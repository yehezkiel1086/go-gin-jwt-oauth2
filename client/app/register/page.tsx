"use client";
import Link from "next/link";
import { useState } from "react";
import { FcGoogle } from "react-icons/fc";

const RegisterPage = () => {
  const [email, setEmail] = useState("");
  const [password, setPassword] = useState("");
  const [passConf, setPassConf] = useState("");
  const [error, setError] = useState("");
  const [loading, setLoading] = useState(false);

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();
    setError("");
    setLoading(true);

    const c1 = email;
    const c2 = password;
    const c3 = passConf == password;

    if (!c1 || !c2 || !c3) {
      setError(
        "Email and password are required. Password confirmation must match password."
      );
      setLoading(false);
      return;
    }

    try {
      const res = await fetch(
        `${process.env.NEXT_PUBLIC_API_URI}/api/v1/register`,
        {
          method: "POST",
          headers: { "Content-Type": "application/json" },
          body: JSON.stringify({ email, password }),
          credentials: "include",
        }
      );

      if (!res.ok) {
        const data = await res.json();
        setError(data.error || "Registration failed");
      } else {
        // Redirect or reload on success
        window.location.href = "/employees";
      }
    } catch (err) {
      setError("Network error");
    } finally {
      setLoading(false);
    }
  };

  return (
    <>
      <h1 className="text-2xl font-semibold">Register</h1>
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
        />
        <div className="flex flex-col gap-2">
          <button
            className="bg-blue-600 rounded-sm py-1 hover:bg-blue-500 transition duration-200"
            type="submit"
            disabled={loading}
          >
            {loading ? "Signing up..." : "Signup"}
          </button>
          <button className="border-blue-600 border-2 rounded-sm py-1 hover:bg-white hover:text-black transition duration-200">
            <FcGoogle className="inline-block text-lg mb-[3px]" /> Signup with
            Google
          </button>
        </div>
        {error && <div className="text-red-600">{error}</div>}
      </form>
      <p className="mt-4">
        Already have an account?{" "}
        <Link href="/login" className="text-blue-600 hover:text-red-600">
          Signin
        </Link>
      </p>
    </>
  );
};

export default RegisterPage;
