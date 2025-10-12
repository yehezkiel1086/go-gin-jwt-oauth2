"use client";
import Link from "next/link";
import { useState } from "react";
import { FcGoogle } from "react-icons/fc";

const Homepage = () => {
  const [email, setEmail] = useState("");
  const [password, setPassword] = useState("");
  const [passConf, setPassConf] = useState("");
  const [error, setError] = useState("");
  const [loading, setLoading] = useState(false);

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();
    setError("");
    setLoading(true);

    if (!email || !password) {
      setError("Email and password are required.");
      setLoading(false);
      return;
    }

    try {
      const res = await fetch("/api/login", {
        method: "POST",
        headers: { "Content-Type": "application/json" },
        body: JSON.stringify({ email, password }),
        credentials: "include",
      });

      if (!res.ok) {
        const data = await res.json();
        setError(data.error || "Login failed");
      } else {
        // Redirect or reload on success
        window.location.href = "/dashboard";
      }
    } catch (err) {
      setError("Network error");
    } finally {
      setLoading(false);
    }
  };

  return (
    <div className="flex items-center justify-center h-screen">
      {/* inner container */}
      <div className="w-96">
        <h1 className="text-2xl font-semibold">Login</h1>
        <form className="flex flex-col gap-4 mt-4" onSubmit={handleSubmit}>
          <input
            type="text"
            name="email"
            id="email"
            placeholder="Enter your email here"
            className="border-1 rounded-sm px-2 py-1"
            value={email}
            onChange={(e) => setEmail(e.target.value)}
          />
          <input
            type="password"
            name="password"
            id="password"
            placeholder="Enter your password here"
            className="border-1 rounded-sm px-2 py-1"
            value={password}
            onChange={(e) => setPassword(e.target.value)}
          />
          <input
            type="password"
            name="pass_conf"
            id="pass_conf"
            placeholder="Re-enter your password here"
            className="border-1 rounded-sm px-2 py-1"
            value={passConf}
            onChange={(e) => setPassConf(e.target.value)}
            disabled
          />
          <div className="flex flex-col gap-2">
            <button
              className="bg-blue-600 rounded-sm py-1 hover:bg-blue-500 transition duration-200"
              type="submit"
              disabled={loading}
            >
              {loading ? "Signing in..." : "Signin"}
            </button>
            <button className="border-blue-600 border-2 rounded-sm py-1 hover:bg-white hover:text-black transition duration-200">
              <FcGoogle className="inline-block text-lg mb-[3px]" /> Login with
              Google
            </button>
          </div>
          {error && <div className="text-red-600">{error}</div>}
        </form>
        <p className="mt-4">
          Don&apos;t have an account?{" "}
          <Link href="/register" className="text-blue-600 hover:text-red-600">
            Sign up
          </Link>
        </p>
      </div>
    </div>
  );
};

export default Homepage;
