"use client";

import Link from "next/link";
import { useActionState, useState } from "react";
import { createEmployee } from "@/app/actions/employee";

const CreateEmployeePage = () => {
  const [name, setName] = useState("");
  const [position, setPosition] = useState("");
  const [description, setDescription] = useState("");

  const [state, action, pending] = useActionState(createEmployee, undefined);

  return (
    <div>
      <Link href="/employees" className="text-blue-600 hover:text-red-600">/employees</Link>
      <h1 className="text-2xl mt-1 font-semibold">Add New Employee</h1>
      <form action={action} className="flex flex-col gap-4 mt-4">
        <input
          type="text"
          name="name"
          id="name"
          placeholder="Enter employee name here"
          className="border-1 rounded-sm px-2 py-1"
          value={name}
          onChange={(e) => setName(e.target.value)}
        />
        {state?.errors?.name && (
          <p className="text-red-500">{state.errors.name}</p>
        )}

        <input
          type="text"
          name="position"
          id="position"
          placeholder="Enter employee position here"
          className="border-1 rounded-sm px-2 py-1"
          value={position}
          onChange={e => setPosition(e.target.value)}
        />
        {state?.errors?.position && (
          <p className="text-red-500">{state.errors.position}</p>
        )}

        <textarea
          name="description" 
          id="description" 
          placeholder="Enter employee description here"
          className="border-1 rounded-sm px-2 py-1"
          value={description}
          onChange={e => setDescription(e.target.value)}
        ></textarea>
        {state?.errors?.description && (
          <p className="text-red-500">{state.errors.description}</p>
        )}

        <button
          className="bg-blue-600 rounded-sm py-1 hover:bg-blue-500 transition duration-200"
          type="submit"
          disabled={pending}
        >
          {pending ? "Adding Employee..." : "Add Employee"}
        </button>
        {/* {error && <div className="text-red-600">{error}</div>} */}
      </form>
    </div>
  );
};

export default CreateEmployeePage;
