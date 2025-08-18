import {useState} from "react";

const useListAllUsers = () => {
  const [loading, setLoading] = useState(false);
  const [error, setError] = useState(null);
  const [users, setUsers] = useState([]);
  const abortController = new AbortController();
  const listAllUsers = async () => {
    setLoading(true);
    const signal = abortController.signal;
    try {
      const res = await fetch("/api/v1/user", {
        method: "GET",
        headers: { "Content-Type": "application/json" },
        signal: signal,
      });

      const data = await res.json();
      if (data.error) {
        throw new Error(data.error)
      }

      setUsers(data.user)
    } catch (err) {
      setError(err)
    } finally {
      setLoading(false)
    }
  };
  
  return {loading, error, users, abortController, listAllUsers}
};

export default useListAllUsers;
