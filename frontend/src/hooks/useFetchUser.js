import {useState} from 'react'

const useFetchUser = () => {
  const [loading, setLoading]=useState(false)
  const [error, setError]=useState(null)
  const [user,setUser]=useState(null)
  const abortController = new AbortController()

  const fetchUser = async (id) => {
    setLoading(true)
    const signal = abortController.signal
    try {
      const res=await fetch(`/api/v1/user/${id}`, {
        method:"GET",
        headers:{"Content-Type":"application/json"},
        signal:signal
      })

      const data=res.json()
      if (data.error) {
        throw new Error(data.error);        
      }

      setUser(data.user)
    } catch (err) {
      setError(err)
    } finally {
      setLoading(false)
    }
  }

  return {loading, error, user, abortController, fetchUser}
}

export default useFetchUser