import {useState} from "react";

const useUpdateUser = () => {
  const { setAuthUser } = useAuthContext();
  const [loading, setLoading] = useState(false);
  const [error, setError] = useState(null);
  const updateUser = async ({ id, name, email, password }) => {
    setLoading(true);
    try {
      const res = await fetch(`/api/v1/user/${id}`, {
        method: "PUT",
        headers: { "Content-Type": "application/json" },
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
  return { loading, error, updateUser };
};

export default useUpdateUser;
