import apiClient from "@/lib/api/client";
import type { LoginInput, RegisterInput } from "@/lib/validators/auth";
import type { UserResponse } from "@/lib/validators/user";

export interface AuthResponse {
  token: string;
  user: UserResponse;
}

export const authService = {
  async login(data: LoginInput): Promise<AuthResponse> {
    const response = await apiClient.post<AuthResponse>("/api/v1/auth/login", data);
    return response.data;
  },

  async register(data: RegisterInput): Promise<AuthResponse> {
    const response = await apiClient.post<AuthResponse>("/api/v1/auth/register", data);
    return response.data;
  },

  async getMe(): Promise<UserResponse> {
    const response = await apiClient.get<UserResponse>("/api/v1/user/me");
    return response.data;
  },

  async updateMe(data: Partial<UserResponse>): Promise<UserResponse> {
    const response = await apiClient.patch<UserResponse>("/api/v1/user/me", data);
    return response.data;
  },
};
