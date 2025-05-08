import {
  SidebarInset,
  SidebarProvider,
  SidebarTrigger,
} from "@/components/ui/Sidebar";
import { Link, Outlet, useLocation } from "react-router";
import { AppSidebar } from "./AppSidebar";
import { Separator } from "@/components/ui/Separator";
import {
  Breadcrumb,
  BreadcrumbItem,
  BreadcrumbLink,
  BreadcrumbList,
  BreadcrumbPage,
  BreadcrumbSeparator,
} from "@/components/ui/Breadcrumb";
import { ModeToggle } from "@/components/theme/ModeToggle";

export function AppLayout() {
  const location = useLocation();
  const paths = location.pathname.split("/").filter((p) => p);

  return (
    <SidebarProvider>
      <AppSidebar />
      <SidebarInset>
        <header className="flex sticky top-0 bg-background h-16 shrink-0 items-center gap-2 border-b px-4">
          <nav className="flex items-center gap-2">
            <SidebarTrigger />
            <Separator orientation="vertical" className="mr-2 h-4" />
            <Breadcrumb>
              <BreadcrumbList>
                <BreadcrumbItem className="hidden md:block">
                  <BreadcrumbLink asChild>
                    <Link to="/app">some dropdown</Link>
                  </BreadcrumbLink>
                </BreadcrumbItem>
                <BreadcrumbSeparator className="hidden md:block" />
                {paths
                  .filter((p) => p !== "app")
                  .map((path) => (
                    <BreadcrumbItem key={path}>
                      <BreadcrumbPage>{path}</BreadcrumbPage>
                    </BreadcrumbItem>
                  ))}
              </BreadcrumbList>
            </Breadcrumb>
          </nav>
          <div className="flex-1" />
          <div className="flex items-center gap-4">
            <ModeToggle />
          </div>
        </header>
        <main className="p-4">
          <Outlet />
        </main>
      </SidebarInset>
    </SidebarProvider>
  );
}
