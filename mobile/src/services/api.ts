const API_BASE_URL = process.env.EXPO_PUBLIC_API_URL || 'http://localhost:8080';

// API Response types
export interface ApiResponse<T> {
  data?: T;
  message?: string;
  error?: string;
}

// User types
export interface User {
  id: string;
  email: string;
  name?: string;
  createdAt: string;
}

export interface LoginRequest {
  email: string;
  password: string;
}

export interface RegisterRequest {
  email: string;
  password: string;
  name?: string;
}

// Workout types
export interface Exercise {
  name: string;
  sets: Set[];
}

export interface Set {
  reps: number;
  weight: number;
  completed: boolean;
}

export interface Workout {
  id?: string;
  userId?: string;
  exercises: Exercise[];
  startedAt?: string;
  completedAt?: string;
  notes?: string;
}

// API Client class
class ApiClient {
  private baseURL: string;
  private authToken: string | null = null;

  constructor(baseURL: string) {
    this.baseURL = baseURL;
  }

  setAuthToken(token: string) {
    this.authToken = token;
  }

  clearAuthToken() {
    this.authToken = null;
  }

  private async request<T>(
    endpoint: string,
    options: RequestInit = {}
  ): Promise<ApiResponse<T>> {
    const url = `${this.baseURL}${endpoint}`;

    const headers: Record<string, string> = {
      'Content-Type': 'application/json',
    };

    if (this.authToken) {
      headers.Authorization = `Bearer ${this.authToken}`;
    }

    try {
      const response = await fetch(url, {
        ...options,
        headers,
      });

      const responseData = await response.json();

      if (!response.ok) {
        return {
          error: responseData.message || 'Request failed',
        };
      }

      return {
        data: responseData,
      };
    } catch (error) {
      return {
        error: error instanceof Error ? error.message : 'Network error',
      };
    }
  }

  // Authentication endpoints
  async register(
    data: RegisterRequest
  ): Promise<ApiResponse<{ user: User; token: string }>> {
    return this.request('/api/auth/register', {
      method: 'POST',
      body: JSON.stringify(data),
    });
  }

  async login(
    data: LoginRequest
  ): Promise<ApiResponse<{ user: User; token: string }>> {
    return this.request('/api/auth/login', {
      method: 'POST',
      body: JSON.stringify(data),
    });
  }

  // User endpoints
  async getUserProfile(): Promise<ApiResponse<User>> {
    return this.request('/api/user/profile');
  }

  async updateUserProfile(data: Partial<User>): Promise<ApiResponse<User>> {
    return this.request('/api/user/profile', {
      method: 'PUT',
      body: JSON.stringify(data),
    });
  }

  // Workout endpoints
  async createWorkout(workout: Workout): Promise<ApiResponse<Workout>> {
    return this.request('/api/workouts', {
      method: 'POST',
      body: JSON.stringify(workout),
    });
  }

  async getWorkouts(): Promise<ApiResponse<Workout[]>> {
    return this.request('/api/workouts');
  }

  async getWorkout(id: string): Promise<ApiResponse<Workout>> {
    return this.request(`/api/workouts/${id}`);
  }

  // Buddy endpoints
  async inviteBuddy(email: string): Promise<ApiResponse<{ message: string }>> {
    return this.request('/api/buddies/invite', {
      method: 'POST',
      body: JSON.stringify({ email }),
    });
  }

  async getBuddies(): Promise<ApiResponse<User[]>> {
    return this.request('/api/buddies');
  }
}

// Export singleton instance
export const apiClient = new ApiClient(API_BASE_URL);

// Health check function
export const checkApiHealth = async (): Promise<boolean> => {
  try {
    const response = await fetch(`${API_BASE_URL}/health`);
    return response.ok;
  } catch {
    return false;
  }
};
