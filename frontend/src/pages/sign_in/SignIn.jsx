const SignIn = () => {
  return (
    <div className="fles flex-col items-center justify-center min-w-96 mx-auto">
      <div className="w-full p-6 rounded-lg shadow-md bg-green-400/10 bg-clip-padding backdrop-filter backdrop-blur-lg">
        <h1 className="text-3xl font-semibold text-center text-gray-300">
          Sign In
        </h1>
        <form>
          <div>
            <label className="label">
              <span className="text-base label-text">Email</span>
            </label>
            <input
              type="email"
              className="w-full input input-bordered-h10"
              placeholder="Email"
            />
          </div>
          <div>
            <label className="label">
              <span className="text-base label-text">Password</span>
            </label>
            <input
              type="password"
              placeholder="Password"
              className="w-full input input-bordered h--10"
            />
          </div>
          <a
            href="#"
            className="text-sm hover:underline hover:text-blue-600 mt-2 inline-block"
          >
            {"Don't"} have an account?
          </a>
          <div>
            <button className="btn btn-block btn-sm mt-2">Sign In</button>
          </div>
        </form>
      </div>
    </div>
  );
};

export default SignIn;
