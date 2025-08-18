import {useState} from "react";
import {useAuthContext} from "../contexts/AuthContext";

const useSignOut = () => {
  const { setAuthUser } = useAuthContext();
  const [loading,setLoading]=useState(false)
  const [error,setError]=useState(null)
  const signout = async () => {
    setLoading(true)
    try {
      const res = await fetch("/api/v1/auth/signout", {
        method: "POST",
        headers: { "Content-Type": "application/json" },
      });

      const data = await res.json();
      if (data.error) {
        throw new Error(data.error);
      }

      localStorage.removeItem("user");
      setAuthUser(null);
    } catch (err) {
      setError(err)
    } finally {
      setLoading(false)
    }
  };
  return {loading, error, signout};
};

export default useSignOut;
