import { useState } from "react";
import ConnectForm from "./components/ConnectForm";
import SchemaList from "./components/SchemaList";
import TableView from "./components/TableView";

function App() {
  const [connected, setConnected] = useState(false);
  const [selectedTable, setSelectedTable] = useState<string | null>(null);

  if (!connected) return <ConnectForm onConnected={() => setConnected(true)} />;

  return (
    <div className="flex gap-4">
      <SchemaList onSelectTable={setSelectedTable} />
      {selectedTable && <TableView table={selectedTable} />}
    </div>
  );
}

export default App;