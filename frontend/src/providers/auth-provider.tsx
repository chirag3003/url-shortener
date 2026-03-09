"use client";

import { authService } from "@/services/auth.service";
import type { UserResponse } from "@/lib/validators/user";
import { useQuery, useQueryClient } from "@tanstack/react-query";
import {
  createContext,
  useCallback,
  useContext,
  useEffect,
  useState,
} from "react";
import { useRouter } from "next/navigation";

interface AuthContextType {
  user: UserResponse | null;
  isLoading: boolean;
  isAuthenticated: boolean;
  login: (token: string, user: UserResponse) => void;
  logout: () => void;
  refreshUser: () => Promise<void>;
}

export const AuthContext = createContext<AuthContextType | undefined>(
  undefined,
);

export function AuthProvider({ children }: { children: React.ReactNode }) {
  const [token, setToken] = useState<string | null>(null);
  const [isInitializing, setIsInitializing] = useState(true);
  const queryClient = useQueryClient();
  const router = useRouter();
  // On mount, check localStorage for token
  useEffect(() => {
    const storedToken = localStorage.getItem("token");
    if (storedToken) {
      setToken(storedToken);
    }
    setIsInitializing(false);
  }, []);

  // Use React Query to manage the user profile state
  const {
    data: user,
    isLoading: queryIsLoading,
    refetch,
  } = useQuery({
    queryKey: ["user", token],
    queryFn: () => authService.getMe(),
    enabled: !!token, // Only fetch if we have a token
    retry: false,
    staleTime: 1000 * 60 * 10, // 10 minutes
  });

  const login = useCallback(
    (newToken: string, newUser: UserResponse) => {
      localStorage.setItem("token", newToken);
      setToken(newToken);
      queryClient.setQueryData(["user", newToken], newUser);
      router.push("/dashboard");
    },
    [queryClient, router],
  );

  const logout = useCallback(() => {
    localStorage.removeItem("token");
    setToken(null);
    queryClient.clear();
    router.push("/");
  }, [queryClient, router]);

  const refreshUser = useCallback(async () => {
    await refetch();
  }, [refetch]);

  const value = {
    user: user?.data ?? null,
    isLoading: isInitializing || (!!token && queryIsLoading),
    isAuthenticated: !!user,
    login,
    logout,
    refreshUser,
  };

  return <AuthContext.Provider value={value}>{children}</AuthContext.Provider>;
}
