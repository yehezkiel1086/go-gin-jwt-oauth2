"use client";

import { useEffect, useState } from "react";
import Link from "next/link";

const EmployeesPage = () => {
  const [employees, setEmployees] = useState<any[]>([]);
  const [loading, setLoading] = useState(true);

  useEffect(() => {
    const getEmployees = async () => {
      try {
        const res = await fetch(
          `${process.env.NEXT_PUBLIC_API_URI}/api/v1/employees`,
          {
            credentials: "include",
          }
        );
        const data = await res.json();
        setEmployees(data);
      } catch (err) {
        console.error("Failed to load employees:", err);
      } finally {
        setLoading(false);
      }
    };
    getEmployees();
  }, []);

  return (
    <div>
      <h1 className="text-2xl font-semibold">Employees Dashboard</h1>
      <Link
        href="/employees/create"
        className="text-blue-600 hover:text-red-600"
      >
        Add new employee
      </Link>

      <div className="mt-4 flex flex-col gap-4">
        {loading ? (
          <p>Loading employees...</p>
        ) : employees?.length === 0 ? (
          <p>No employees found.</p>
        ) : (
          employees?.map((emp, i) => (
            <div key={i}>
              <Link
                href={`/employees/${emp.id}`}
                className="text-lg text-blue-600 hover:text-red-600"
              >
                {emp.name}
              </Link>
              <p>Position: {emp.position}</p>
            </div>
          ))
        )}
      </div>
    </div>
  );
};

export default EmployeesPage;
