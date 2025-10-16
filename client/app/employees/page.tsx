import Link from "next/link"

const EmployeesPage = () => {
  const employees = [
    {
      "id": 1,
      "name": "Dimitry Kravchenko",
      "position": "Soviet Army",
    },
    {
      "id": 2,
      "name": "Imran Zhakaev",
      "position": "Russian Ultranationalist",
    },
    {
      "id": 3,
      "name": "John MacTavish",
      "position": "SAS",
    }
  ]

  return (
    <div>
      <h1 className="text-2xl font-semibold">Employees Dashboard</h1>
      <Link
        href="/employees/create"
        className="text-blue-600 hover:text-red-600"
      >
        Add new employee
      </Link>
      {/* employees listing */}
      <div className="mt-4 flex flex-col gap-4">
        {employees.map(
          (emp, i) =>
            i < 10 && (
              <div key={emp.id}>
                <Link
                  href={`/employees/${emp.id}`}
                  className="text-lg text-blue-600 hover:text-red-600"
                >
                  {emp.name}
                </Link>
                <p>Position: {emp.position}</p>
              </div>
            )
        )}
      </div>
    </div>
  );
}

export default EmployeesPage
