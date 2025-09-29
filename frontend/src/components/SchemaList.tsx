import { useState, useEffect } from "react";
import axios from "axios";

interface Props {
  onSelectTable: (table: string) => void;
}

export default function SchemaList({ onSelectTable }: Props) {
  const [tables, setTables] = useState<any[]>([]);

  useEffect(() => {
    axios.get("http://localhost:8080/schema").then((res) => setTables(res.data));
  }, []);

  return (
    <div className="p-4 border rounded w-60">
      <h2 className="font-bold mb-2">Tables</h2>
      <ul>
        {tables.map((t) => (
          <li
            key={t.name}
            className="cursor-pointer hover:underline"
            onClick={() => onSelectTable(t.name)}
          >
            {t.name}
          </li>
        ))}
      </ul>
    </div>
  );
}