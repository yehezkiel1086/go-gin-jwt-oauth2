import Footer from "@/app/components/Footer";

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
        <Footer />
      </div>
    </div>
  );
};

export default EmployeesLayout;
