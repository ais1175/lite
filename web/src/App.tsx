import { QueryClient, QueryClientProvider } from "@tanstack/react-query";
import { BrowserRouter, Route, Routes } from "react-router";
import { AuthRoute } from "./features/auth/routes/AuthRoute";
import { AppDashboard } from "./features/app/routes/AppDashboard";
import { AppLayout } from "./features/app/components/AppLayout";
import { TokensRoute } from "./features/tokens/routes/TokensRoute";
import { ThemeProvider } from "./components/theme/ThemeProvider";
import { ProtectedRoute } from "./features/auth/routes/ProtectedRoute";
import { NewOrganizationRoute } from "./features/organizations/routes/NewOrganizationRoute";

const queryClient = new QueryClient();

function App() {
  return (
    <QueryClientProvider client={queryClient}>
      <ThemeProvider>
        <BrowserRouter>
          <Routes>
            <Route path="*" element={<div>404</div>} />
            <Route path="/auth" element={<AuthRoute />} />
            <Route path="/app">
              <Route
                path="new-organization"
                element={<NewOrganizationRoute />}
              />
              <Route element={<ProtectedRoute />}>
                <Route path=":organizationId" element={<AppLayout />}>
                  <Route index element={<AppDashboard />} />
                  <Route path="tokens" element={<TokensRoute />} />
                </Route>
              </Route>
            </Route>
          </Routes>
        </BrowserRouter>
      </ThemeProvider>
    </QueryClientProvider>
  );
}

export default App;
