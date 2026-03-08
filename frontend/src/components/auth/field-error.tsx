"use client";

export function FieldError({ message }: { message?: string }) {
  if (!message) return null;
  return (
    <p className="text-xs text-destructive/90 font-medium animate-in fade-in slide-in-from-top-1 duration-150 mt-1.5">
      {message}
    </p>
  );
}
