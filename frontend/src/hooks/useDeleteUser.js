import {useState} from "react";
import {useAuthContext} from "../contexts/AuthContext";

const useDeleteUser = () => {
  const { setAuthUser } = useAuthContext();
  const [loading, setLoading] = useState(false);
  const [error, setError] = useState(null);
  const deleteUser = async (id) => {
    setLoading(true);
    try {
      const res = await fetch(`/api/v1/user/${id}`, {
        method: "DELETE",
        headers: { "Content-type": "application/json" },
      });

      localStorage.removeItem("user");
      setAuthUser(null);
    } catch (err) {
      setError(err);
    } finally {
      setLoading(false);
    }
  };
  return { loading, error, deleteUser };
};

export default useDeleteUser;
