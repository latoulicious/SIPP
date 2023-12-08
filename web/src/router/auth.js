export default function (to, from, next) {
  const jwtToken = localStorage.getItem("jwtToken");

  // Check if the user is authenticated (you might want to add more checks)
  if (!jwtToken) {
    // Redirect to the login page if not authenticated
    next({ name: "login" }); // Use the name of the route
  } else {
    // Proceed to the next route if authenticated
    next();
  }
}
