import { useEffect, useState } from "react";
import axios from "axios";

interface Props {
  table: string;
}

export default function TableView({ table }: Props) {
  const [rows, setRows] = useState<any[]>([]);

  useEffect(() => {
    axios
      .get(`http://localhost:8080/table/${table}?limit=10&offset=0`)
      .then((res) => setRows(res.data));
  }, [table]);

  if (!rows.length) return <div>No data</div>;

  const cols = Object.keys(rows[0]);

  return (
    <div className="overflow-x-auto p-4">
      <h2 className="font-bold mb-2">Table: {table}</h2>
      <table className="border-collapse border border-gray-400">
        <thead>
          <tr>
            {cols.map((c) => (
              <th key={c} className="border p-1">
                {c}
              </th>
            ))}
          </tr>
        </thead>
        <tbody>
          {rows.map((r, i) => (
            <tr key={i}>
              {cols.map((c) => (
                <td key={c} className="border p-1">
                  {String(r[c])}
                </td>
              ))}
            </tr>
          ))}
        </tbody>
      </table>
    </div>
  );
}