import EditProfile from "../../components/user/EditProfile";

const Home = () => {
  return (
    <div className="flex rounded-lg overflow-hidden bg-green-400/10 bg-clip-padding backdrop-filter backdrop-blur-lg sm:h-[450px] md:h-[550px]">
      <EditProfile />
    </div>
  );
};

export default Home;
