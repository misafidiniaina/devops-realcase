import React, { useEffect, useState } from 'react'
import { createRoot } from 'react-dom/client'
import './style.css'

const API = import.meta.env.VITE_API_URL || 'http://localhost:8080'
const empty = {
  name: '',
  hostname: '',
  ip_address: '',
  environment: 'development',
  provider: 'aws',
}

function App() {
  const [servers, setServers] = useState([])
  const [form, setForm] = useState(empty)
  const [editingId, setEditingId] = useState(null)
  const [error, setError] = useState('')
  const [loading, setLoading] = useState(true)

  async function load() {
    setLoading(true)
    try {
      const response = await fetch(`${API}/api/v1/servers`)
      if (!response.ok) throw Error('Failed to load servers')
      setServers(await response.json())
      setError('')
    } catch (e) {
      setError(e.message)
    } finally {
      setLoading(false)
    }
  }

  useEffect(() => {
    load()
  }, [])

  function change(e) {
    setForm({ ...form, [e.target.name]: e.target.value })
  }

  function startEdit(server) {
    setEditingId(server.id)
    setForm({
      name: server.name,
      hostname: server.hostname,
      ip_address: server.ip_address,
      environment: server.environment,
      provider: server.provider,
    })
    setError('')
  }

  function cancelEdit() {
    setEditingId(null)
    setForm(empty)
    setError('')
  }

  async function submit(e) {
    e.preventDefault()
    const endpoint = editingId
      ? `${API}/api/v1/servers/${editingId}`
      : `${API}/api/v1/servers`
    const method = editingId ? 'PUT' : 'POST'

    try {
      const response = await fetch(endpoint, {
        method,
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify(form),
      })
      const body = await response.json()
      if (!response.ok) throw Error(body.error || 'Failed to save server')
      cancelEdit()
      await load()
    } catch (e) {
      setError(e.message)
    }
  }

  async function remove(id) {
    if (!confirm('Delete this server?')) return
    const response = await fetch(`${API}/api/v1/servers/${id}`, { method: 'DELETE' })
    if (!response.ok) setError('Failed to delete server')
    else load()
  }

  return (
    <main>
      <header>
        <p className="eyebrow">CLOUD PLATFORM LAB · PHASE 1</p>
        <h1>Infrastructure Inventory</h1>
        <p className="intro">A small production-style workload for the DevOps and DevSecOps platform.</p>
      </header>

      <section className="panel">
        <h2>{editingId ? 'Edit server' : 'Add server'}</h2>
        <form onSubmit={submit}>
          <input name="name" placeholder="Name" value={form.name} onChange={change} required />
          <input name="hostname" placeholder="Hostname" value={form.hostname} onChange={change} required />
          <input name="ip_address" placeholder="IP address" value={form.ip_address} onChange={change} required />
          <select name="environment" value={form.environment} onChange={change}>
            <option>development</option>
            <option>staging</option>
            <option>production</option>
          </select>
          <select name="provider" value={form.provider} onChange={change}>
            <option>aws</option>
            <option>gcp</option>
            <option>on-premise</option>
          </select>
          <button type="submit">{editingId ? 'Save changes' : 'Create server'}</button>
          {editingId && <button type="button" className="secondary" onClick={cancelEdit}>Cancel</button>}
        </form>
      </section>

      {error && <p className="error">{error}</p>}

      <section className="panel">
        <div className="section-heading">
          <h2>Servers</h2>
          <button className="secondary" onClick={load}>Refresh</button>
        </div>
        {loading ? <p>Loading...</p> : servers.length === 0 ? <p className="muted">No infrastructure registered yet.</p> :
          <div className="table-wrap">
            <table>
              <thead><tr><th>Name</th><th>Hostname</th><th>IP</th><th>Environment</th><th>Provider</th><th></th></tr></thead>
              <tbody>{servers.map((server) => <tr key={server.id}>
                <td>{server.name}</td>
                <td>{server.hostname}</td>
                <td>{server.ip_address}</td>
                <td>{server.environment}</td>
                <td>{server.provider}</td>
                <td><div className="actions"><button className="secondary" onClick={() => startEdit(server)}>Edit</button><button className="danger" onClick={() => remove(server.id)}>Delete</button></div></td>
              </tr>)}</tbody>
            </table>
          </div>}
      </section>
    </main>
  )
}

createRoot(document.getElementById('root')).render(<App />)
