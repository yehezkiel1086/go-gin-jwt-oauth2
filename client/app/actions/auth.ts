import { SignupFormSchema, FormState } from '@/app/lib/definitions'
 
export async function signup(state: FormState, formData: FormData) {
  // Validate form fields
  const validatedFields = SignupFormSchema.safeParse({
    name: formData.get('name'),
    email: formData.get('email'),
    password: formData.get('password'),
    pass_conf: formData.get('pass_conf'),
  })
 
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
          name: formData.get('name'),
          email: formData.get('email'),
          password: formData.get('password'),
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
          pass_conf: data.error?.pass_conf ? [data.error.pass_conf] : undefined,
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
        pass_conf: ["Network error"],
      },
    };
  }
}
