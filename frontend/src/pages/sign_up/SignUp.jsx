const SignUp = () => {
  return (
    <div className="flex flex-col items-center justify-center min-w-96 mx-auto">
      <div className="w-full p-6 rounded-lg shadow-md bg-green-400/10 bg-clip-padding backdrop-filter backdrop-blur-lg">
        <h1 className="text-3xl font-semibold text-center text-gray-300">
          Sign Up
        </h1>
        <form>
          <div>
            <label className="label">
              <span className="text-base label-text">Full Name</span>
            </label>
            <input
              type="text"
              placeholder="Full Name"
              className="w-full input input-bordered h-10"
            />
          </div>
          <div>
            <label className="label">
              <span className="text-base label-text">Email</span>
            </label>
            <input
              type="email"
              placeholder="Email"
              className="w-full input input-bordered h-10"
            />
          </div>
          <div>
            <label className="label">Password</label>
            <input
              type="password"
              placeholder="Password"
              className="w-full input input-bordered h-10"
            />
          </div>
          <a
            href=""
            className="text-sm hover:underline hover:text-blue-600 mt-2 inline-block"
          >
            Already have an account?
          </a>
          <div>
            <button
              className="btn btn-block btn-sm mt-2"
              onClick={(e) =>{
                e.preventDefault()
                document.getElementById("sign_up_modal").showModal()}
              }
            >
              Sign Up
            </button>
          </div>
        </form>
        <dialog id="sign_up_modal" className="modal">
          <div className="modal-box">
            <h3 className="font-bold text-lg">New Account! || Error!</h3>
            <p className="py-4">Message</p>
            <div className="modal-action">
              <form method="dialog">
                <button className="btn btn-block btn-sm mt-2">
                  Sign In || Close
                </button>
              </form>
            </div>
          </div>
        </dialog>
      </div>
    </div>
  );
};

export default SignUp;
