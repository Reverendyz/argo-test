import { useState } from 'react';
import './App.css';
import Button from '@mui/material/Button';

const HOST = process.env.API_HOST || '127.0.0.1';
const PORT = process.env.API_PORT || '8080';

function App() {
  const [response, setResponse] = useState(null);

  const fetchTime = async () => {
    const endpoint = `http://${HOST}:${PORT}/timer`;
    try {
      const res = await fetch(endpoint);
      const data = await res.json();
      setResponse(data);
    } catch (error) {
      alert('Failed to fetch data. Check the endpoint.');
    }
  };

  return (
    <div className="App">
      <Button variant="contained" color="success" onClick={fetchTime}>
        Fetch Time from Endpoint
      </Button>
      {response && (
        <div style={{ marginTop: '20px' }}>
          <p>
            <strong>Hostname:</strong> {response.hostname}
          </p>
          <p>
            <strong>Date:</strong> {response.time}
          </p>
        </div>
      )}
    </div>
  );
}

export default App;
