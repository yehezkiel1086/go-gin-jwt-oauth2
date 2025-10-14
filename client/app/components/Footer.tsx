"use client";

import { useActionState } from "react";
import { logout } from "@/app/actions/auth";

const Footer = () => {
  const [state, action, pending] = useActionState(logout, undefined);

  return (
    <footer className="mt-4 text-right block">
      <form action={action}>
        <button
          className="text-blue-600 hover:text-red-600"
          disabled={pending}
          type="submit"
        >
          {pending ? "Logging out..." : "Logout"}
        </button>
        {state?.error && <p className="text-red-500 mt-2">{state.error}</p>}
      </form>
    </footer>
  );
}

export default Footer
