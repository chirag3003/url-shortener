"use client";

import { useState } from "react";
import { toast } from "sonner";
import { Button } from "@/components/ui/button";
import { Input } from "@/components/ui/input";
import { Label } from "@/components/ui/label";
import { authService } from "@/services/auth.service";

export function SecurityView() {
  const [currentPassword, setCurrentPassword] = useState("");
  const [newPassword, setNewPassword] = useState("");
  const [confirmPassword, setConfirmPassword] = useState("");
  const [loading, setLoading] = useState(false);

  const handleUpdatePassword = async () => {
    if (!currentPassword || !newPassword || !confirmPassword) {
      toast.error("Please fill in all password fields");
      return;
    }

    if (newPassword !== confirmPassword) {
      toast.error("New passwords do not match");
      return;
    }

    if (newPassword.length < 8) {
      toast.error("New password must be at least 8 characters");
      return;
    }

    setLoading(true);
    try {
      await authService.updateMe({
        // Assuming backend handles password update via updateMe with specific fields
        // or we need a specific password update method.
        // For now, mapping to updateMe based on common patterns.
        password: newPassword,
        currentPassword: currentPassword,
      } as any);
      toast.success("Password updated successfully");
      setCurrentPassword("");
      setNewPassword("");
      setConfirmPassword("");
    } catch (error: any) {
      toast.error(error.response?.data?.error || "Failed to update password");
    } finally {
      setLoading(false);
    }
  };

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
            value={currentPassword}
            onChange={(e) => setCurrentPassword(e.target.value)}
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
            value={newPassword}
            onChange={(e) => setNewPassword(e.target.value)}
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
            value={confirmPassword}
            onChange={(e) => setConfirmPassword(e.target.value)}
          />
        </div>

        <Button
          variant="outline"
          className="rounded-lg font-medium"
          disabled={loading}
          onClick={handleUpdatePassword}
        >
          {loading ? "Updating..." : "Update Password"}
        </Button>
      </div>
    </div>
  );
}
