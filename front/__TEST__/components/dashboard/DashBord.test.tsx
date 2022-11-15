import { screen } from "@testing-library/react";
import DashBord from "@/components/dashboard/DashBord";
import "@testing-library/jest-dom";
import { renderWithProviders } from "../../utils/test-utils";
import { server } from "../../../src/mocks/server";
import { rest } from "msw";

const BACKEND_API_URL: string = process.env.NEXT_PUBLIC_BACKEND_API_URL!;

describe("DashBord", () => {
  it("render DashBord", async () => {
    server.use(
      rest.get(BACKEND_API_URL + "/dashboard", (req, res, ctx) => {
        // If authenticated, return a mocked user details
        return res(
          ctx.status(200),
          ctx.json({
            siteSize: 57,
            articleSize: 6810,
          })
        );
      })
    );

    renderWithProviders(<DashBord />);

    const loading = screen.getByText(/Loading/i);
    expect(loading).toBeInTheDocument();

    const site = await screen.findByText(/Site: 57/i);
    expect(site).toBeInTheDocument();

    const article = await screen.findByText(/Article: 6810/i);
    expect(article).toBeInTheDocument();
  });
});
