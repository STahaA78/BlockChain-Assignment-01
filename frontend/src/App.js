import React, { useState, useEffect } from "react";
import axios from "axios";
import './index.css'; // Make sure CSS is imported

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
    if (!transaction.trim()) return;
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
    if (!search.trim()) return;
    const res = await axios.get(`${API}/search?data=${search}`);
    setSearchResult(res.data);
  };

  useEffect(() => {
    fetchChain();
    fetchPending();
  }, []);

  return (
    <div className="app-container">
      <header className="app-header">
        {/* Required Assignment Header */}
        <h1>Syed_Taha_Ahmed Blockchain</h1>
      </header>

      <main className="main-content">
        {/* Actions Panel */}
        <section className="card action-panel">
          <div className="input-group">
            <h2>Add Transaction</h2>
            <div className="flex-row">
              <input
                type="text"
                placeholder="Enter transaction data..."
                value={transaction}
                onChange={(e) => setTransaction(e.target.value)}
              />
              <button className="btn-primary" onClick={addTransaction}>Add</button>
            </div>
          </div>

          <div className="input-group">
            <h2>Search Transaction</h2>
            <div className="flex-row">
              <input
                type="text"
                placeholder="Search by data (e.g., i22-6638)..."
                value={search}
                onChange={(e) => setSearch(e.target.value)}
              />
              <button className="btn-secondary" onClick={searchData}>Search</button>
            </div>
          </div>

          <div className="mine-group">
            <h2>Mine Pending Block</h2>
            <button className="btn-mine" onClick={mineBlock}>⛏️ Mine Block</button>
          </div>
        </section>

        {/* Data Display Panel */}
        <section className="data-panel">
          <div className="card">
            <h2>Pending Transactions</h2>
            {pending.length === 0 ? (
              <p className="empty-state">No pending transactions.</p>
            ) : (
              <pre className="json-display">{JSON.stringify(pending, null, 2)}</pre>
            )}
          </div>

          {searchResult && (
            <div className="card highlight-card">
              <h2>Search Result</h2>
              <pre className="json-display">{JSON.stringify(searchResult, null, 2)}</pre>
            </div>
          )}

          <div className="card">
            <h2>Blockchain Ledger</h2>
            <pre className="json-display blockchain-display">
              {JSON.stringify(chain, null, 2)}
            </pre>
          </div>
        </section>
      </main>
    </div>
  );
}

export default App;