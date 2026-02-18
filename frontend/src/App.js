import React, { useState, useEffect } from "react";
import axios from "axios";

const API = "http://localhost:8080";

function App() {
  const [transaction, setTransaction] = useState("");
  const [chain, setChain] = useState([]);
  const [pending, setPending] = useState([]);
  const [search, setSearch] = useState("");
  const [searchResult, setSearchResult] = useState(null);

  const fetchChain = async () => {
    const res = await axios.get(`${API}/chain`);
    setChain(res.data);
  };

  const fetchPending = async () => {
    const res = await axios.get(`${API}/pending`);
    setPending(res.data);
  };

  const addTransaction = async () => {
    await axios.post(`${API}/add`, { transaction });
    setTransaction("");
    fetchPending();
  };

  const mineBlock = async () => {
    await axios.post(`${API}/mine`);
    fetchChain();
    fetchPending();
  };

  const searchData = async () => {
    const res = await axios.get(`${API}/search?data=${search}`);
    setSearchResult(res.data);
  };

  useEffect(() => {
    fetchChain();
    fetchPending();
  }, []);

  return (
    <div style={{ padding: 20 }}>
      <h1>Syed_Taha_Ahmed Blockchain</h1>

      <h2>Add Transaction</h2>
      <input
        value={transaction}
        onChange={(e) => setTransaction(e.target.value)}
      />
      <button onClick={addTransaction}>Add</button>

      <h2>Pending Transactions</h2>
      <pre>{JSON.stringify(pending, null, 2)}</pre>

      <h2>Mine Block</h2>
      <button onClick={mineBlock}>Mine</button>

      <h2>Blockchain</h2>
      <pre>{JSON.stringify(chain, null, 2)}</pre>

      <h2>Search Transaction</h2>
      <input value={search} onChange={(e) => setSearch(e.target.value)} />
      <button onClick={searchData}>Search</button>
      <pre>{JSON.stringify(searchResult, null, 2)}</pre>
    </div>
  );
}

export default App;
