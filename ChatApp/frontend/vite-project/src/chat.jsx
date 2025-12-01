import { useState } from 'react';


export default function Chat() {
const [message, setMessage] = useState('');
const [history, setHistory] = useState([]);
const [loading, setLoading] = useState(false);


async function send() {
if (!message) return;
const userMsg = { role: 'user', content: message };
setHistory(h => [...h, { from: 'you', text: message }]);
setMessage('');
setLoading(true);


try {
const res = await fetch('http://localhost:8080/chat', {
method: 'POST',
headers: { 'Content-Type': 'application/json' },
body: JSON.stringify({ message: message }),
});


const data = await res.json();
if (res.ok) {
setHistory(h => [...h, { from: 'bot', text: data.reply }]);
} else {
setHistory(h => [...h, { from: 'bot', text: 'Error: ' + (data.error || 'unknown') }]);
}
} catch (err) {
setHistory(h => [...h, { from: 'bot', text: 'Network error: ' + err.message }]);
} finally {
setLoading(false);
}
}


function onKey(e) {
if (e.key === 'Enter') send();
}


return (
<div className="chat-container">
<div className="history">
{history.map((m, i) => (
<div key={i} className={"msg " + m.from}>
<strong>{m.from === 'you' ? 'You' : 'Bot'}:</strong>
<div>{m.text}</div>
</div>
))}
</div>


<div className="input-row">
<input
value={message}
onChange={e => setMessage(e.target.value)}
onKeyDown={onKey}
placeholder="Type a message and press Enter or Send"
type="text"
/>
<button onClick={send} disabled={loading}>{loading ? '...' : 'Send'}</button>
</div>
</div>
);
}