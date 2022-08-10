package main

import (
	"context"
	"fmt"

	// "fmt"
	"os"
	// "os/signal"
	// "sync"
	// "time"

	// "fmt"
	// "os"

	"github.com/jedib0t/go-pretty/table"
	log "github.com/sirupsen/logrus"

	"github.com/gopcua/opcua"
	"github.com/gopcua/opcua/id"

	// "github.com/gopcua/opcua/monitor"
	"github.com/gopcua/opcua/ua"
)

//Read data from OPC UA
// func main() {
// 	endpoint := "opc.tcp://milo.digitalpetri.com:62541/milo"
// 	nodeID := "ns=2;s=Dynamic/RandomFloat"

// 	ctx := context.Background()

// 	c := opcua.NewClient(endpoint, opcua.SecurityMode(ua.MessageSecurityModeNone))
// 	if err := c.Connect(ctx); err != nil {
// 		log.Fatal(err)
// 	}
// 	defer c.Close()

// 	id, err := ua.ParseNodeID(nodeID)
// 	if err != nil {
// 		log.Fatalf("invalid node id: %v", err)
// 	}
// 	for i := 0; i < 10; i++ {

// 		startTime := time.Now()

// 		log.Info("Request started at: ", startTime)

// 		req := &ua.ReadRequest{
// 			MaxAge:             2000,
// 			NodesToRead:        []*ua.ReadValueID{{NodeID: id}},
// 			TimestampsToReturn: ua.TimestampsToReturnBoth,
// 		}

// 		resp, err := c.Read(req)
// 		if err != nil {
// 			log.Fatalf("Read failed: %s", err)
// 		}

// 		if resp.Results[0].Status != ua.StatusOK {
// 			log.Fatalf("Status not OK: %v", resp.Results[0].Status)
// 		}

// 		endTime := time.Now()

// 		log.Info("Response get: ", endTime)

// 		log.Info(endTime.Sub(startTime))

// 		log.Printf("Value: %#v", resp.Results[0].Value.Value())
// 		fmt.Println()
// 	}
// }

// Browse available variables
func getDataType(value *ua.DataValue) string {
	if value.Status != ua.StatusOK {
		return value.Status.Error()
	}

	switch value.Value.NodeID().IntID() {
	case id.DateTime:
		return "time.Time"

	case id.Boolean:
		return "bool"

	case id.Int32:
		return "int32"
	}

	return value.Value.NodeID().String()
}
func main() {
	endpoint := "opc.tcp://milo.digitalpetri.com:62541/milo"
	nodeID := "ns=0;i=85"

	ctx := context.Background()
	// Connect
	c := opcua.NewClient(endpoint)
	if err := c.Connect(ctx); err != nil {
		log.Fatal(err)
	}
	defer c.Close()

	eps, err := c.GetEndpoints()
	if err != nil {
		log.Fatal(err)
	}

	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.AppendHeader(table.Row{"Endpoint", "Policy", "Mode", "Tokens"})
	for _, ep := range eps.Endpoints {
		var tokens []string
		for _, token := range ep.UserIdentityTokens {
			tokens = append(tokens, token.TokenType.String())
		}
		t.AppendRow(table.Row{
			ep.EndpointURL, ep.SecurityPolicyURI, ep.SecurityMode, tokens,
		})

	}
	t.Render()

	nid, err := ua.ParseNodeID(nodeID)
	if err != nil {
		log.Fatalf("invalid node id: %s", nodeID)
	}

	n := c.Node(nid)
	
	attrs, err := n.Attributes(ua.AttributeIDBrowseName, ua.AttributeIDDataType)
	if err != nil {
		log.Fatalf("invalid attribute: %s", err.Error())
	}

	fmt.Printf("BrowseName: %s; DataType: %s\n", attrs[0].Value, getDataType(attrs[1]))

	// Get children
	refs, err := n.ReferencedNodes(id.HasComponent, ua.BrowseDirectionForward, ua.NodeClassAll, true)
	if err != nil {
		log.Fatalf("References: %s", err)
	}

	fmt.Printf("Children: %d\n", len(refs))

	for _, rn := range refs { // Node IDs and Description
		fmt.Printf("Node ID:   %s\n", rn.ID.String())

		desc, err := rn.Description()
		if err != nil {
			log.Error(err)
		}

		if desc.Text != "" {
			fmt.Printf("Node Desc:   %s\n", desc.Text)
		}

		childs, err := rn.Children(rn.ID.IntID(), ua.NodeClassAll) // Node Childrens ID and description
		if err != nil {
			log.Error(err)
		}
		for _, rn := range childs {
			fmt.Printf("  ->Children ID:    %s\n", rn.ID.String())
			desc, err := rn.Description()
			if err != nil {
				log.Error(err)
			}

			if desc.Text != "" {
				fmt.Printf("    Description:   %s\n", desc.Text)
			}
		}
	}
}

// //Subscribe to a OPC Channel
// func cleanup(sub *monitor.Subscription, wg *sync.WaitGroup) {
// 	log.Printf("stats: sub=%d delivered=%d dropped=%d", sub.SubscriptionID(), sub.Delivered(), sub.Dropped())
// 	sub.Unsubscribe(context.Background())
// 	wg.Done()
// }

// func startCallbackSub(ctx context.Context, m *monitor.NodeMonitor, interval, lag time.Duration, wg *sync.WaitGroup, nodes ...string) {
// 	sub, err := m.Subscribe(
// 		ctx,
// 		&opcua.SubscriptionParameters{
// 			Interval: interval,
// 		},
// 		func(s *monitor.Subscription, msg *monitor.DataChangeMessage) {
// 			if msg.Error != nil {
// 				log.Printf("[callback] error=%s", msg.Error)
// 			} else {
// 				log.Printf("[callback] node=%s value=%v", msg.NodeID, msg.Value.Value())
// 			}
// 			time.Sleep(lag)
// 		},
// 		nodes...)

// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	defer cleanup(sub, wg)

// 	<-ctx.Done()
// }

// func main() {
// 	endpoint := "opc.tcp://milo.digitalpetri.com:62541/milo"
// 	nodeID := "ns=2;s=Dynamic/RandomFloat"

// 	signalCh := make(chan os.Signal, 1)
// 	signal.Notify(signalCh, os.Interrupt)

// 	ctx, cancel := context.WithCancel(context.Background())
// 	defer cancel()

// 	go func() {
// 		<-signalCh
// 		println()
// 		cancel()
// 	}()

// 	c := opcua.NewClient(endpoint, opcua.SecurityMode(ua.MessageSecurityModeNone))
// 	if err := c.Connect(ctx); err != nil {
// 		log.Fatal(err)
// 	}
// 	defer c.Close()

// 	m, err := monitor.NewNodeMonitor(c)
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	m.SetErrorHandler(func(_ *opcua.Client, sub *monitor.Subscription, err error) {
// 		log.Printf("error: sub=%d err=%s", sub.SubscriptionID(), err.Error())
// 	})
// 	wg := &sync.WaitGroup{}

// 	// start callback-based subscription
// 	wg.Add(1)
// 	go startCallbackSub(ctx, m, time.Second, 0, wg, nodeID)

// 	<-ctx.Done()
// 	wg.Wait()
// }
