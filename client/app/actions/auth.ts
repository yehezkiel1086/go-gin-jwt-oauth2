import { SignupFormSchema, FormState } from '@/app/lib/definitions'
 
export async function signup(state: FormState, formData: FormData) {
  // 1. Validate form fields
  const validatedFields = SignupFormSchema.safeParse({
    name: formData.get('name'),
    email: formData.get('email'),
    password: formData.get('password'),
  })
 
  // If any form fields are invalid, return early
  if (!validatedFields.success) {
    return {
      errors: validatedFields.error.flatten().fieldErrors,
    }
  }
 
  // 2. Prepare data for insertion into database
  const { name, email, password } = validatedFields.data
 
  // 3. Insert the user into the database or call an Auth Library's API
  const res = await fetch(`${process.env.NEXT_PUBLIC_API_URI}/register`, {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json',
    },
    body: JSON.stringify({
      name,
      email,
      password,
    }),
  })

  const jsonData = await res.json()

  console.log(jsonData);
 
  // TODO:
  // 4. Create user session
  // 5. Redirect user
}