import Link from "next/link";

const Homepage = () => {
  return (
    <div>
      <h1>Homepage</h1>
      <Link href="/register">Register</Link>
    </div>
  );
};

export default Homepage;
