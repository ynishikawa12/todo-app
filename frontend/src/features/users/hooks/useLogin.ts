import { useState } from "react";

export const useLogin = () => {
  const [email, setEmail] = useState("");
  const [password, setPassword] = useState("");
  const [loading, setLoading] = useState(false);
  const [error, setError] = useState<string | null>(null);

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();
    setLoading(true);
    setError(null);

    // TODO: try catch api
  };

  return {
    email,
    setEmail,
    password,
    setPassword,
    loading,
    error,
    handleSubmit,
  };
};
