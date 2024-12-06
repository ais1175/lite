export const IS_DEV = import.meta.env.DEV;

export class ApiError extends Error {
  constructor(error: string, errorMessage: string) {
    super(error);

    this.message = errorMessage;
  }
}

export async function fetchApi(input: string | URL, init?: RequestInit) {
  try {
    const res = await fetch(input, {
      ...init,
      credentials: "include",
    });
    if (!res.ok) {
      const errorResponse = (await res.json()) as {
        error: string;
        message: string;
      };

      throw new ApiError(errorResponse.error, errorResponse.message);
    }

    return await res.json();
  } catch (err) {
    if (err instanceof Error) {
      throw new Error(err.message);
    }
  }
}
