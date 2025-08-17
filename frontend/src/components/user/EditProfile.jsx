const EditProfile = () => {
  return (
    <div className="w-full p-6">
      <h1 className="text-3xl font-semibold text-center text-gray-300">
        Edit Profile
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
        <div>
          <button
            className="btn btn-block btn-sm mt-2"
            onClick={(e) => {
              e.preventDefault();
            }}
          >
            Submit
          </button>
        </div>
      </form>
    </div>
  );
};

export default EditProfile;
