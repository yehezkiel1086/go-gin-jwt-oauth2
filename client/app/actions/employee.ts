import {
  EmployeeFormSchema,
  FormState,
} from "@/app/lib/definitions";

export async function createEmployee(state: FormState, formData: FormData) {
  // Validate form fields
  const validatedFields = EmployeeFormSchema.safeParse({
    name: formData.get("name"),
    position: formData.get("position"),
    description: formData.get("description"),
  });

  // If any form fields are invalid, return early
  if (!validatedFields.success) {
    return {
      errors: validatedFields.error.flatten().fieldErrors,
    };
  }

  const { name, position, description } = validatedFields.data;

  // Call the provider or db to create a user...
  try {
    const res = await fetch(
      `${process.env.NEXT_PUBLIC_API_URI}/api/v1/admin/employees`,
      {
        method: "POST",
        headers: { "Content-Type": "application/json" },
        body: JSON.stringify({
          name, position, description
        }),
        credentials: "include",
      }
    );

    if (!res.ok) {
      const data = await res.json();
      return {
        errors: {
          name: data.error?.name ? [data.error.name] : undefined,
          position: data.error?.position ? [data.error.position] : undefined,
          description: data.error?.description ? [data.error.description] : undefined,
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
