"use client";

import { useState } from "react";
import { toast } from "sonner";
import { Button } from "@/components/ui/button";
import {
  Dialog,
  DialogContent,
  DialogDescription,
  DialogTitle,
  DialogTrigger,
} from "@/components/ui/dialog";
import { Input } from "@/components/ui/input";
import { Label } from "@/components/ui/label";
import type { LoginInput } from "@/lib/validators/auth";
import { loginSchema, registerSchema } from "@/lib/validators/auth";
import { FieldError } from "./auth/field-error";
import { OAuthButtons } from "./auth/oauth-buttons";
import { PasswordStrengthBar } from "./auth/password-strength";

// ── Auth Modal ───────────────────────────────────────────────────────────

export function AuthModal({
  children,
  defaultTab = "login",
}: {
  children: React.ReactNode;
  defaultTab?: "login" | "register";
}) {
  const [open, setOpen] = useState(false);
  const [mode, setMode] = useState<"login" | "register">(defaultTab);

  const [loginValues, setLoginValues] = useState<LoginInput>({
    email: "",
    password: "",
  });
  const [regValues, setRegValues] = useState({
    name: "",
    email: "",
    password: "",
    confirmPassword: "",
  });
  const [errors, setErrors] = useState<Record<string, string>>({});
  const [loading, setLoading] = useState(false);

  const handleOpenChange = (isOpen: boolean) => {
    setOpen(isOpen);
    if (!isOpen) {
      setTimeout(() => {
        setMode(defaultTab);
        setErrors({});
        setLoginValues({ email: "", password: "" });
        setRegValues({
          name: "",
          email: "",
          password: "",
          confirmPassword: "",
        });
      }, 300);
    }
  };

  const handleLoginSubmit = async (e: React.FormEvent) => {
    e.preventDefault();
    setErrors({});
    const parsed = loginSchema.safeParse(loginValues);
    if (!parsed.success) {
      const fieldErrors: Record<string, string> = {};
      for (const issue of parsed.error.issues) {
        fieldErrors[issue.path.join(".")] = issue.message;
      }
      setErrors(fieldErrors);
      return;
    }
    setLoading(true);
    await new Promise((r) => setTimeout(r, 600));
    setLoading(false);
    toast.success("Welcome back!");
    handleOpenChange(false);
  };

  const handleRegSubmit = async (e: React.FormEvent) => {
    e.preventDefault();
    setErrors({});
    const parsed = registerSchema.safeParse(regValues);
    if (!parsed.success) {
      const fieldErrors: Record<string, string> = {};
      for (const issue of parsed.error.issues) {
        fieldErrors[issue.path.join(".")] = issue.message;
      }
      setErrors(fieldErrors);
      return;
    }
    setLoading(true);
    await new Promise((r) => setTimeout(r, 800));
    setLoading(false);
    toast.success("Account created! Welcome aboard.");
    handleOpenChange(false);
  };

  return (
    <Dialog open={open} onOpenChange={handleOpenChange}>
      <DialogTrigger asChild>{children}</DialogTrigger>
      <DialogContent className="sm:max-w-md p-0 overflow-hidden border-border/50 shadow-2xl bg-background/95 backdrop-blur-xl">
        <div className="flex flex-col">
          {/* Header Area */}
          <div className="px-8 pt-8 pb-6 text-center">
            <div className="mx-auto mb-4 flex h-12 w-12 items-center justify-center rounded-2xl bg-primary/10 text-primary shadow-sm ring-1 ring-primary/20">
              <svg
                xmlns="http://www.w3.org/2000/svg"
                width="24"
                height="24"
                viewBox="0 0 24 24"
                fill="none"
                stroke="currentColor"
                strokeWidth="2.5"
                strokeLinecap="round"
                strokeLinejoin="round"
              >
                <path d="M10 13a5 5 0 0 0 7.54.54l3-3a5 5 0 0 0-7.07-7.07l-1.72 1.71" />
                <path d="M14 11a5 5 0 0 0-7.54-.54l-3 3a5 5 0 0 0 7.07 7.07l1.71-1.71" />
              </svg>
            </div>
            <DialogTitle className="text-2xl font-bold tracking-tight">
              {mode === "login" ? "Welcome back" : "Create an account"}
            </DialogTitle>
            <DialogDescription className="mt-2 text-sm text-muted-foreground/80">
              {mode === "login"
                ? "Enter your credentials to access your account"
                : "Enter your details below to create your account"}
            </DialogDescription>
          </div>

          {/* Form Area */}
          <div className="px-8 pb-8">
            {mode === "login" ? (
              <form
                onSubmit={handleLoginSubmit}
                className="space-y-4 animate-in slide-in-from-left-4 fade-in duration-300"
              >
                <div className="space-y-1.5">
                  <Label
                    htmlFor="login-email"
                    className="text-muted-foreground font-medium text-xs uppercase tracking-widest pl-1"
                  >
                    Email
                  </Label>
                  <Input
                    id="login-email"
                    type="email"
                    placeholder="name@example.com"
                    className="h-11 rounded-xl bg-card border-border/50 shadow-sm focus-visible:ring-primary/50"
                    value={loginValues.email}
                    onChange={(e) =>
                      setLoginValues((v) => ({ ...v, email: e.target.value }))
                    }
                  />
                  <FieldError message={errors.email} />
                </div>

                <div className="space-y-1.5">
                  <div className="flex items-center justify-between pl-1">
                    <Label
                      htmlFor="login-password"
                      className="text-muted-foreground font-medium text-xs uppercase tracking-widest"
                    >
                      Password
                    </Label>
                    <button
                      type="button"
                      className="text-xs font-medium text-primary hover:underline hover:text-primary/80"
                    >
                      Forgot?
                    </button>
                  </div>
                  <Input
                    id="login-password"
                    type="password"
                    placeholder="••••••••"
                    className="h-11 rounded-xl bg-card border-border/50 shadow-sm focus-visible:ring-primary/50"
                    value={loginValues.password}
                    onChange={(e) =>
                      setLoginValues((v) => ({
                        ...v,
                        password: e.target.value,
                      }))
                    }
                  />
                  <FieldError message={errors.password} />
                </div>

                <Button
                  type="submit"
                  className="w-full h-11 mt-2 rounded-xl font-semibold shadow-md"
                  disabled={loading}
                >
                  {loading ? "Signing in..." : "Sign In"}
                </Button>
              </form>
            ) : (
              <form
                onSubmit={handleRegSubmit}
                className="space-y-4 animate-in slide-in-from-right-4 fade-in duration-300 max-h-[50vh] overflow-y-auto px-1 -mx-1 pb-1"
              >
                <div className="space-y-1.5">
                  <Label
                    htmlFor="reg-name"
                    className="text-muted-foreground font-medium text-xs uppercase tracking-widest pl-1"
                  >
                    Full Name
                  </Label>
                  <Input
                    id="reg-name"
                    placeholder="John Doe"
                    className="h-11 rounded-xl bg-card border-border/50 shadow-sm focus-visible:ring-primary/50"
                    value={regValues.name}
                    onChange={(e) =>
                      setRegValues((v) => ({ ...v, name: e.target.value }))
                    }
                  />
                  <FieldError message={errors.name} />
                </div>

                <div className="space-y-1.5">
                  <Label
                    htmlFor="reg-email"
                    className="text-muted-foreground font-medium text-xs uppercase tracking-widest pl-1"
                  >
                    Email
                  </Label>
                  <Input
                    id="reg-email"
                    type="email"
                    placeholder="name@example.com"
                    className="h-11 rounded-xl bg-card border-border/50 shadow-sm focus-visible:ring-primary/50"
                    value={regValues.email}
                    onChange={(e) =>
                      setRegValues((v) => ({ ...v, email: e.target.value }))
                    }
                  />
                  <FieldError message={errors.email} />
                </div>

                <div className="space-y-1.5">
                  <Label
                    htmlFor="reg-password"
                    className="text-muted-foreground font-medium text-xs uppercase tracking-widest pl-1"
                  >
                    Password
                  </Label>
                  <Input
                    id="reg-password"
                    type="password"
                    placeholder="••••••••"
                    className="h-11 rounded-xl bg-card border-border/50 shadow-sm focus-visible:ring-primary/50"
                    value={regValues.password}
                    onChange={(e) =>
                      setRegValues((v) => ({ ...v, password: e.target.value }))
                    }
                  />
                  <PasswordStrengthBar password={regValues.password} />
                  <FieldError message={errors.password} />
                </div>

                <div className="space-y-1.5">
                  <Label
                    htmlFor="reg-confirm"
                    className="text-muted-foreground font-medium text-xs uppercase tracking-widest pl-1"
                  >
                    Confirm Password
                  </Label>
                  <Input
                    id="reg-confirm"
                    type="password"
                    placeholder="••••••••"
                    className="h-11 rounded-xl bg-card border-border/50 shadow-sm focus-visible:ring-primary/50"
                    value={regValues.confirmPassword}
                    onChange={(e) =>
                      setRegValues((v) => ({
                        ...v,
                        confirmPassword: e.target.value,
                      }))
                    }
                  />
                  <FieldError message={errors.confirmPassword} />
                </div>

                <Button
                  type="submit"
                  className="w-full h-11 mt-2 rounded-xl font-semibold shadow-md"
                  disabled={loading}
                >
                  {loading ? "Creating account..." : "Sign Up"}
                </Button>
              </form>
            )}

            <OAuthButtons />

            <div className="mt-8 text-center text-sm text-muted-foreground">
              {mode === "login" ? (
                <>
                  Don&apos;t have an account?{" "}
                  <button
                    type="button"
                    onClick={() => setMode("register")}
                    className="font-semibold text-primary hover:underline"
                  >
                    Sign up
                  </button>
                </>
              ) : (
                <>
                  Already have an account?{" "}
                  <button
                    type="button"
                    onClick={() => setMode("login")}
                    className="font-semibold text-primary hover:underline"
                  >
                    Sign in
                  </button>
                </>
              )}
            </div>
          </div>
        </div>
      </DialogContent>
    </Dialog>
  );
}
