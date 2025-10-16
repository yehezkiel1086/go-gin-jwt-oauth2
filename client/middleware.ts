import { NextRequest, NextResponse } from "next/server";
import { cookies } from "next/headers";

// 1. Define public and protected routes
const protectedRoutes = ["/employees"];
const publicRoutes = ["/login", "/register", "/"];

export default async function middleware(req: NextRequest) {
  // 2. Check if the current route is protected or public
  const path = req.nextUrl.pathname;
  const isProtectedRoute = protectedRoutes.includes(path);
  const isPublicRoute = publicRoutes.includes(path);

  // 3. Decrypt the session from the cookie
  const cookie = (await cookies()).get("jwt_token")?.value;

  // 4. Redirect to /login if the user is not authenticated
  if (isProtectedRoute && !cookie) {
    return NextResponse.redirect(new URL("/login", req.nextUrl));
  }

  // 5. Redirect to /dashboard if the user is authenticated
  if (
    isPublicRoute &&
    cookie &&
    !req.nextUrl.pathname.startsWith("/employees")
  ) {
    return NextResponse.redirect(new URL("/employees", req.nextUrl));
  }

  return NextResponse.next();
}

// Routes Middleware should not run on
export const config = {
  matcher: ["/((?!api|_next/static|_next/image|.*\\.png$).*)"],
};
