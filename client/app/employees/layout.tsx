import Link from "next/link";

const EmployeesLayout = ({
  children,
}: Readonly<{
  children: React.ReactNode;
}>) => {
  return (
    <div className="flex items-center justify-center h-screen">
      {/* inner container */}
      <div className="w-96">
        {children}
        <Link href="/logout" className="text-blue-600 hover:text-red-600 mt-4 text-right block">Logout</Link>
      </div>
    </div>
  )
}

export default EmployeesLayout;
