"use client";

import Link from "next/link";
import { useParams } from "next/navigation";
import { useEffect, useState } from "react";

// Dummy data for demonstration
const employees = [
  {
    id: 1,
    name: "Dimitry Kravchenko",
    position: "Soviet Army",
  },
  {
    id: 2,
    name: "Imran Zhakaev",
    position: "Russian Ultranationalist",
  },
  {
    id: 3,
    name: "John MacTavish",
    position: "SAS",
  },
];

const EmployeePage = () => {
  const params = useParams();
  const id = params.id ? Number(params.id) : null;

  const [name, setName] = useState("");
  const [position, setPosition] = useState("");
  const [error, setError] = useState("");
  const [loading, setLoading] = useState(false);

  useEffect(() => {
    if (id !== null) {
      const emp = employees.find((emp) => emp.id === id);
      if (emp) {
        setName(emp.name);
        setPosition(emp.position);
      } else {
        setError("Employee not found.");
      }
    }
  }, [id]);

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();
    setError("");
    setLoading(true);

    if (!name || !position) {
      setError("Name and position are required.");
      setLoading(false);
      return;
    }

    // Example: Replace with actual update logic
    setTimeout(() => {
      setLoading(false);
      alert("Employee updated (dummy logic)");
    }, 1000);
  };

  return (
    <div>
      <Link href="/employees" className="text-blue-600 hover:text-red-600">/employees</Link>
      <h1 className="text-2xl mt-1 font-semibold">Edit Employee</h1>
      <form onSubmit={handleSubmit} className="flex flex-col gap-4 mt-4">
        <input
          type="text"
          name="name"
          id="name"
          placeholder="Enter employee name here"
          className="border-1 rounded-sm px-2 py-1"
          value={name}
          onChange={(e) => setName(e.target.value)}
        />
        <input
          type="text"
          name="position"
          id="position"
          placeholder="Enter position here"
          className="border-1 rounded-sm px-2 py-1"
          value={position}
          onChange={e => setPosition(e.target.value)}
        />
        <button
          className="bg-blue-600 rounded-sm py-1 hover:bg-blue-500 transition duration-200"
          type="submit"
          disabled={loading}
        >
          {loading ? "Saving..." : "Save"}
        </button>
        {error && <div className="text-red-600">{error}</div>}
      </form>
    </div>
  );
};

export default EmployeePage;
