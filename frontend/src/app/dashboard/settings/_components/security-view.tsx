"use client";

import { toast } from "sonner";
import { Button } from "@/components/ui/button";
import { Input } from "@/components/ui/input";
import { Label } from "@/components/ui/label";

export function SecurityView() {
  return (
    <div className="space-y-6 animate-in fade-in duration-300">
      <div>
        <h2 className="text-xl font-semibold tracking-tight">Security</h2>
        <p className="text-sm text-muted-foreground mt-1">
          Manage your password and account security preferences.
        </p>
      </div>
      <div className="h-px bg-border/60" />

      <div className="space-y-8 max-w-xl">
        <div className="space-y-3">
          <Label htmlFor="current-pass" className="text-sm font-medium">
            Current Password
          </Label>
          <Input
            id="current-pass"
            type="password"
            placeholder="••••••••••••"
            className="max-w-md h-10"
          />
        </div>

        <div className="space-y-3">
          <Label htmlFor="new-pass" className="text-sm font-medium">
            New Password
          </Label>
          <Input
            id="new-pass"
            type="password"
            placeholder="••••••••••••"
            className="max-w-md h-10"
          />
          <p className="text-[13px] text-muted-foreground">
            Must be at least 8 characters long.
          </p>
        </div>

        <div className="space-y-3">
          <Label htmlFor="confirm-pass" className="text-sm font-medium">
            Confirm New Password
          </Label>
          <Input
            id="confirm-pass"
            type="password"
            placeholder="••••••••••••"
            className="max-w-md h-10"
          />
        </div>

        <Button
          variant="outline"
          className="rounded-lg font-medium"
          onClick={() => toast.success("Password changed successfully")}
        >
          Update Password
        </Button>
      </div>
    </div>
  );
}
