"use client";
import Link from "next/link";
import { useState, useActionState } from "react";
import { FcGoogle } from "react-icons/fc";
import { signup } from "@/app/actions/auth";

const RegisterPage = () => {
  const [name, setName] = useState("");
  const [email, setEmail] = useState("");
  const [password, setPassword] = useState("");
  const [passConf, setPassConf] = useState("");

  const [state, action, pending] = useActionState(signup, undefined);

  return (
    <>
      <h1 className="text-2xl font-semibold">Register</h1>
      <form className="flex flex-col gap-4 mt-4" action={action}>
        <input
          type="text"
          name="name"
          id="name"
          placeholder="Enter your name here"
          className="border-1 rounded-sm px-2 py-1"
          value={name}
          onChange={(e) => setName(e.target.value)}
        />
        {state?.errors?.name && (
          <p className="text-red-500">{state.errors.name}</p>
        )}

        <input
          type="email"
          name="email"
          id="email"
          placeholder="Enter your email here"
          className="border-1 rounded-sm px-2 py-1"
          value={email}
          onChange={(e) => setEmail(e.target.value)}
        />
        {state?.errors?.email && (
          <p className="text-red-500">{state.errors.email}</p>
        )}

        <input
          type="password"
          name="password"
          id="password"
          placeholder="Enter your password here"
          className="border-1 rounded-sm px-2 py-1"
          value={password}
          onChange={(e) => setPassword(e.target.value)}
        />
        {state?.errors?.password && (
          <p className="text-red-500">{state.errors.password}</p>
        )}

        <input
          type="password"
          name="password_confirmation"
          id="password_confirmation"
          placeholder="Re-enter your password here"
          className="border-1 rounded-sm px-2 py-1"
          value={passConf}
          onChange={(e) => setPassConf(e.target.value)}
        />
        {state?.errors?.password_confirmation && (
          <p className="text-red-500">{state.errors.password_confirmation}</p>
        )}

        <div className="flex flex-col gap-2">
          <button
            className="bg-blue-600 rounded-sm py-1 hover:bg-blue-500 transition duration-200"
            type="submit"
            disabled={pending}
          >
            {pending ? "Signing up..." : "Signup"}
          </button>
          <button className="border-blue-600 border-2 rounded-sm py-1 hover:bg-white hover:text-black transition duration-200">
            <FcGoogle className="inline-block text-lg mb-[3px]" /> Signup with
            Google
          </button>
        </div>
        {/* {error && <div className="text-red-600">{error}</div>} */}
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
