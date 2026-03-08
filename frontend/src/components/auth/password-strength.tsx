"use client";

import { cn } from "@/lib/utils";

export function getPasswordStrength(password: string): {
  score: number;
  label: string;
  color: string;
} {
  let score = 0;
  if (password.length >= 8) score++;
  if (password.length >= 12) score++;
  if (/[A-Z]/.test(password)) score++;
  if (/[0-9]/.test(password)) score++;
  if (/[^a-zA-Z0-9]/.test(password)) score++;

  if (score <= 1) return { score, label: "Weak", color: "bg-red-500" };
  if (score <= 2) return { score, label: "Fair", color: "bg-orange-500" };
  if (score <= 3) return { score, label: "Good", color: "bg-yellow-500" };
  return { score, label: "Strong", color: "bg-green-500" };
}

export function PasswordStrengthBar({ password }: { password: string }) {
  const { score, label, color } = getPasswordStrength(password);
  if (!password) return null;

  return (
    <div className="space-y-1.5 animate-in fade-in duration-200 mt-2">
      <div className="flex h-1 gap-1 rounded-full overflow-hidden">
        {[1, 2, 3, 4, 5].map((i) => (
          <div
            key={i}
            className={cn(
              "h-full flex-1 rounded-full transition-colors duration-300",
              i <= score ? color : "bg-muted",
            )}
          />
        ))}
      </div>
      <p className="text-[10px] text-muted-foreground font-medium uppercase tracking-wider">
        {label}
      </p>
    </div>
  );
}
