const Profile = () => {
  return (
    <div className="card w-96 shadow-sm">
      <div className="card-body">
        <h2 className="text-3xl font-semibold md:text-left text-center text-gray-300">
          Profile
        </h2>
        <ul className="list bg-base-100/10 rounded-box shadow-md">
          <li className="list-row">
            <div>
              <img
                src="https://img.daisyui.com/images/profile/demo/1@94.webp"
                className="size-10 rounded-box"
              />
            </div>
            <div className="overflow-hidden">
              <span className="block text-nowrap">Jhon Doe</span>
              <span className="text-xs lowercase font-semibold opacity-60 block text-nowrap">
                Jhon@doe.co.us
              </span>
            </div>
            <button
              type="button"
              title="Edit"
              id="edit"
              className="btn btn-square btn-ghost"
            >
              <svg
                className="size-[1.2em]"
                viewBox="0 0 24 24"
                fill="none"
                xmlns="http://www.w3.org/2000/svg"
              >
                <g id="SVGRepo_bgCarrier" stroke-width="0"></g>
                <g
                  id="SVGRepo_tracerCarrier"
                  stroke-linecap="round"
                  stroke-linejoin="round"
                ></g>
                <g id="SVGRepo_iconCarrier">
                  {" "}
                  <path
                    d="M12 3.99997H6C4.89543 3.99997 4 4.8954 4 5.99997V18C4 19.1045 4.89543 20 6 20H18C19.1046 20 20 19.1045 20 18V12M18.4142 8.41417L19.5 7.32842C20.281 6.54737 20.281 5.28104 19.5 4.5C18.7189 3.71895 17.4526 3.71895 16.6715 4.50001L15.5858 5.58575M18.4142 8.41417L12.3779 14.4505C12.0987 14.7297 11.7431 14.9201 11.356 14.9975L8.41422 15.5858L9.00257 12.6441C9.08001 12.2569 9.27032 11.9013 9.54951 11.6221L15.5858 5.58575M18.4142 8.41417L15.5858 5.58575"
                    stroke="currentColor"
                    stroke-width="2"
                    stroke-linecap="round"
                    stroke-linejoin="round"
                  ></path>{" "}
                </g>
              </svg>
            </button>
            <button
              type="button"
              title="Delete"
              id="delete"
              className="btn btn-square btn-ghost"
              onClick={(e) => {
                e.preventDefault();
                document.getElementById("delete_modal").showModal();
              }}
            >
              <svg
                className="size-[1.2em]"
                viewBox="0 0 24 24"
                fill="none"
                xmlns="http://www.w3.org/2000/svg"
                stroke="currentColor"
              >
                <g id="SVGRepo_bgCarrier" stroke-width="0"></g>
                <g
                  id="SVGRepo_tracerCarrier"
                  stroke-linecap="round"
                  stroke-linejoin="round"
                ></g>
                <g id="SVGRepo_iconCarrier">
                  {" "}
                  <path
                    d="M10 12V17"
                    stroke="currentColor"
                    stroke-width="2"
                    stroke-linecap="round"
                    stroke-linejoin="round"
                  ></path>{" "}
                  <path
                    d="M14 12V17"
                    stroke="currentColor"
                    stroke-width="2"
                    stroke-linecap="round"
                    stroke-linejoin="round"
                  ></path>{" "}
                  <path
                    d="M4 7H20"
                    stroke="currentColor"
                    stroke-width="2"
                    stroke-linecap="round"
                    stroke-linejoin="round"
                  ></path>{" "}
                  <path
                    d="M6 10V18C6 19.6569 7.34315 21 9 21H15C16.6569 21 18 19.6569 18 18V10"
                    stroke="currentColor"
                    stroke-width="2"
                    stroke-linecap="round"
                    stroke-linejoin="round"
                  ></path>{" "}
                  <path
                    d="M9 5C9 3.89543 9.89543 3 11 3H13C14.1046 3 15 3.89543 15 5V7H9V5Z"
                    stroke="currentColor"
                    stroke-width="2"
                    stroke-linecap="round"
                    stroke-linejoin="round"
                  ></path>{" "}
                </g>
              </svg>
            </button>
          </li>
          <li className="p-4 pb-2 text-xs opacity-60 tracking-wide">
            {new Date().toLocaleString()}
          </li>
        </ul>
      </div>
      <dialog id="delete_modal" className="modal">
        <div className="modal-box">
          <h3 className="font-bold text-lg">Delete Account</h3>
          <p className="py-4">Confirm to delete your account</p>
          <div className="modal-action">
            <form method="dialog" className="flex flex-row justify-around items-end">
              <button
                title="Cancel Delete"
                className="btn btn-outline btn-sm mt-2 mx-2"
              >
                Cancel
              </button>
              <button
                title="Confirm Delete"
                className="btn btn-outline btn-sm mt-2 mx-2 btn-secondary"
              >
                Confirm
              </button>
            </form>
          </div>
        </div>
      </dialog>
    </div>
  );
};

export default Profile;
