package main

import (
	"context"
	"fmt"
	"time"

	// "fmt"
	// "os"

	log "github.com/sirupsen/logrus"
	// "github.com/jedib0t/go-pretty/table"

	"github.com/gopcua/opcua"
	// "github.com/gopcua/opcua/id"
	"github.com/gopcua/opcua/ua"
)

func main() {
	endpoint := "opc.tcp://milo.digitalpetri.com:62541/milo"
	nodeID := "ns=2;s=Dynamic/RandomFloat"

	ctx := context.Background()

	c := opcua.NewClient(endpoint, opcua.SecurityMode(ua.MessageSecurityModeNone))
	if err := c.Connect(ctx); err != nil {
		log.Fatal(err)
	}
	defer c.Close()

	id, err := ua.ParseNodeID(nodeID)
	if err != nil {
		log.Fatalf("invalid node id: %v", err)
	}
	for i := 0; i < 10; i++ {

		startTime := time.Now()

		log.Info("Request started at: ", startTime)

		req := &ua.ReadRequest{
			MaxAge:             2000,
			NodesToRead:        []*ua.ReadValueID{{NodeID: id}},
			TimestampsToReturn: ua.TimestampsToReturnBoth,
		}

		resp, err := c.Read(req)
		if err != nil {
			log.Fatalf("Read failed: %s", err)
		}

		if resp.Results[0].Status != ua.StatusOK {
			log.Fatalf("Status not OK: %v", resp.Results[0].Status)
		}

		endTime := time.Now()

		log.Info("Response get: ", endTime)

		log.Info(endTime.Sub(startTime))

		log.Printf("Value: %#v", resp.Results[0].Value.Value())
		fmt.Println()
	}
}

// func getDataType(value *ua.DataValue) string {
// 	if value.Status != ua.StatusOK {
// 		return value.Status.Error()
// 	}

// 	switch value.Value.NodeID().IntID() {
// 	case id.DateTime:
// 		return "time.Time"

// 	case id.Boolean:
// 		return "bool"

// 	case id.Int32:
// 		return "int32"
// 	}

// 	return value.Value.NodeID().String()
// }
// func main() {
// 	endpoint := "opc.tcp://milo.digitalpetri.com:62541/milo"
// 	nodeID := "ns=0;i=85"

// 	ctx := context.Background()
// 	// Connect
// 	c := opcua.NewClient(endpoint)
// 	if err := c.Connect(ctx); err != nil {
// 		log.Fatal(err)
// 	}
// 	defer c.Close()

// 	eps, err := c.GetEndpoints()
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	t := table.NewWriter()
// 	t.SetOutputMirror(os.Stdout)
// 	t.AppendHeader(table.Row{"Endpoint", "Policy", "Mode", "Tokens"})
// 	for _, ep := range eps.Endpoints {
// 		var tokens []string
// 		for _, token := range ep.UserIdentityTokens {
// 			tokens = append(tokens, token.TokenType.String())
// 		}
// 		t.AppendRow(table.Row{
// 			ep.EndpointURL, ep.SecurityPolicyURI, ep.SecurityMode, tokens,
// 		})

// 	}
// 	t.Render()

// 	nid, err := ua.ParseNodeID(nodeID)
// 	if err != nil {
// 		log.Fatalf("invalid node id: %s", nodeID)
// 	}

// 	n := c.Node(nid)
// 	attrs, err := n.Attributes(ua.AttributeIDBrowseName, ua.AttributeIDDataType)
// 	if err != nil {
// 		log.Fatalf("invalid attribute: %s", err.Error())
// 	}

// 	fmt.Printf("BrowseName: %s; DataType: %s\n", attrs[0].Value, getDataType(attrs[1]))

// 	// Get children
// 	refs, err := n.ReferencedNodes(id.HasComponent, ua.BrowseDirectionForward, ua.NodeClassAll, true)
// 	if err != nil {
// 		log.Fatalf("References: %s", err)
// 	}

// 	fmt.Printf("Children: %d\n", len(refs))
// 	for _, rn := range refs {
// 		fmt.Printf("   %s\n", rn.ID.String())
// 	}
// }
