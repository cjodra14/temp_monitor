// import logo from './logo.svg';
import './App.css';
import Header from "./components/Header"
import axios from 'axios'


var statusData 

axios.get("http://localhost:8080/status")
    .then(response => {
        statusData = "Temperature "+response.data.Temperature+"ºC and humidity "+response.data.Humidity+"%"
        console.log(statusData)
    })
    .catch(error => {
        console.log(error)
    })

function App() {
  return (
    <div className="App">
      <Header/>
      <h2>{statusData}</h2>

      {/* <header className="App-header">
        <img src={logo} className="App-logo" alt="logo" />
        <p>
          Edit <code>src/App.js</code> and save to reload.
        </p>
        <a
          className="App-link"
          href="https://reactjs.org"
          target="_blank"
          rel="noopener noreferrer"
        >
          Learn React
        </a>
      </header> */}

    </div>
  );
}

export default App;
