import { useEffect, useState } from "react";

function App() {
  const [message, setMessage] = useState("Loading...");

  useEffect(() => {
    fetch("http://localhost:8080/api/message")
      .then((res) => res.json())
      .then((data) => setMessage(data.message))
      .catch((err) => console.error("Error fetching data:", err));
  }, []);

  return (
    <div>
      <h1>Frontend</h1>
      <p>Backend says: {message}</p>
    </div>
  );
}

export default App;
