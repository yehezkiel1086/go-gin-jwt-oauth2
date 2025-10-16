import { z } from 'zod'
 
export const SignupFormSchema = z
  .object({
    name: z
      .string()
      .min(2, { message: "Name must be at least 2 characters long." })
      .trim(),
    email: z.string().email({ message: "Please enter a valid email." }).trim(),
    password: z
      .string()
      .min(8, { message: "Be at least 8 characters long" })
      .regex(/[a-zA-Z]/, { message: "Contain at least one letter." })
      .regex(/[0-9]/, { message: "Contain at least one number." })
      .regex(/[^a-zA-Z0-9]/, {
        message: "Contain at least one special character.",
      })
      .trim(),
    password_confirmation: z.string().trim(),
  })
  .refine((data) => data.password === data.password_confirmation, {
    message: "Password confirmation must match password.",
    path: ["password_confirmation"],
  });

// --- Add this for signin ---
export const SigninFormSchema = z.object({
  email: z.string().email({ message: "Please enter a valid email." }).trim(),
  password: z
    .string()
    .min(8, { message: "At least 8 characters password is required." })
    .trim(),
});

// --- Employee Schema ---
export const EmployeeFormSchema = z.object({
  name: z
    .string()
    .min(2, { message: "Must be at least 2 characters." })
    .max(255, { message: "Name is too long." })
    .trim(),
  position: z
    .string()
    .min(2, { message: "Must be at least 2 characters." })
    .max(255, { message: "Position is too long." })
    .trim(),
  description: z
    .string()
    .max(255, { message: "Description is too long." })
    .optional()
});

export type FormState =
  | {
      errors?: {
        name?: string[];
        email?: string[];
        password?: string[];
        password_confirmation?: string[];
        position?: string[];
        description?: string[];
      };
      message?: string;
    }
  | undefined;
