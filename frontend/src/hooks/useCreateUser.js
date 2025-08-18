import {useState} from "react";
import {useAuthContext} from "../contexts/AuthContext";

const useCreateUser = () => {
  const { setAuthUser } = useAuthContext();
  const [loading, setLoading] = useState(false);
  const [error, setError] = useState(null);
  const createUser = async ({ name, email, password }) => {
    setLoading(true);
    try {
      const res = await fetch("/api/v1/user", {
        method: "POST",
        headers: { "Content-Type": "appliction/json" },
        body: JSON.stringify({ name, email, password }),
      });

      const data = await res.json();
      if (data.error) {
        throw new Error(data.error);
      }

      localStorage.setItem("user", JSON.stringify(data.user));
      setAuthUser(data);
    } catch (err) {
      setError(err);
    } finally {
      setLoading(false);
    }
  };
  return { loading, error, createUser };
};

export default useCreateUser;
