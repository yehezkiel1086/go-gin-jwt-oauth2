import {
  SigninFormSchema,
  SignupFormSchema,
  FormState,
} from "@/app/lib/definitions";

export async function signup(state: FormState, formData: FormData) {
  // Validate form fields
  const validatedFields = SignupFormSchema.safeParse({
    name: formData.get("name"),
    email: formData.get("email"),
    password: formData.get("password"),
    password_confirmation: formData.get("password_confirmation"),
  });

  // If any form fields are invalid, return early
  if (!validatedFields.success) {
    return {
      errors: validatedFields.error.flatten().fieldErrors,
    };
  }

  // Call the provider or db to create a user...
  try {
    const res = await fetch(
      `${process.env.NEXT_PUBLIC_API_URI}/api/v1/register`,
      {
        method: "POST",
        headers: { "Content-Type": "application/json" },
        body: JSON.stringify({
          name: formData.get("name"),
          email: formData.get("email"),
          password: formData.get("password"),
        }),
        credentials: "include",
      }
    );

    if (!res.ok) {
      const data = await res.json();
      return {
        errors: {
          name: data.error?.name ? [data.error.name] : undefined,
          email: data.error?.email ? [data.error.email] : undefined,
          password: data.error?.password ? [data.error.password] : undefined,
          password_confirmation: data.error?.password_confirmation
            ? [data.error.password_confirmation]
            : undefined,
        },
      };
    } else {
      // Redirect or reload on success
      window.location.href = "/employees";
      return {};
    }
  } catch (err) {
    return {
      errors: {
        password_confirmation: ["Network error"],
      },
    };
  }
}

export async function signin(state: FormState, formData: FormData) {
  // Validate form fields
  const validatedFields = SigninFormSchema.safeParse({
    email: formData.get("email"),
    password: formData.get("password"),
  });

  // If any form fields are invalid, return early
  if (!validatedFields.success) {
    return {
      errors: validatedFields.error.flatten().fieldErrors,
    };
  }

  try {
    const res = await fetch(`${process.env.NEXT_PUBLIC_API_URI}/api/v1/login`, {
      method: "POST",
      headers: { "Content-Type": "application/json" },
      body: JSON.stringify({
        email: formData.get("email"),
        password: formData.get("password"),
      }),
      credentials: "include",
    });

    if (!res.ok) {
      const data = await res.json();
      return {
        errors: {
          email: data.error?.email ? [data.error.email] : undefined,
          password: data.error ? [data.error] : undefined,
        },
      };
    } else {
      // Redirect or reload on success
      window.location.href = "/employees";
      return {};
    }
  } catch (err) {
    return {
      errors: {
        password: ["Network error"],
      },
    };
  }
}

export async function logout() {
  try {
    const res = await fetch(
      `${process.env.NEXT_PUBLIC_API_URI}/api/v1/logout`,
      {
        method: "POST",
        credentials: "include",
      }
    );

    if (!res.ok) {
      const data = await res.json();
      console.error("Logout failed:", data);
      return { error: data.message || "Logout failed" };
    }

    // redirect to login page or homepage
    window.location.href = "/login";
    return {};
  } catch (err) {
    return { error: "Network error" };
  }
}
