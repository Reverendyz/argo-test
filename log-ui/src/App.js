import { useState } from 'react';
import './App.css';
import Button from '@mui/material/Button';

// Use environment variables with a default value if not set
const HOST = import.meta.env.VITE_API_HOST || '0.0.0.0';
const PORT = import.meta.env.VITE_API_PORT || '3000';

function App() {
  const [response, setResponse] = useState(null);

  const fetchTime = async () => {
    const endpoint = `http://${HOST}:${PORT}/time`;

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
            <strong>Time:</strong> {response.time}
          </p>
        </div>
      )}
    </div>
  );
}

export default App;
