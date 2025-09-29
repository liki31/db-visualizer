import { useState } from "react";
import axios from "axios";

interface Props {
  onConnected: () => void;
}

export default function ConnectForm({ onConnected }: Props) {
  const [form, setForm] = useState({
    host: "localhost",
    port: 5432,
    user: "postgres",
    password: "secret",
    database: "testdb",
  });

  const handleChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    setForm({ ...form, [e.target.name]: e.target.value });
  };

  const connect = async () => {
    try {
      await axios.post("http://localhost:8080/connect", form);
      onConnected();
    } catch (err: any) {
      alert("Failed to connect: " + err.message);
    }
  };

  return (
    <div className="p-4 border rounded w-80">
      <h2 className="text-lg font-bold mb-2">Connect to Database</h2>
      {Object.keys(form).map((key) => (
        <div key={key} className="mb-2">
          <input
            className="border p-1 w-full"
            placeholder={key}
            name={key}
            type={key === "password" ? "password" : "text"}
            value={(form as any)[key]}
            onChange={handleChange}
          />
        </div>
      ))}
      <button
        onClick={connect}
        className="bg-blue-500 text-white px-3 py-1 rounded"
      >
        Connect
      </button>
    </div>
  );
}