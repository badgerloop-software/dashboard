import { GitDashPage } from './app.po';

describe('git-dash App', () => {
  let page: GitDashPage;

  beforeEach(() => {
    page = new GitDashPage();
  });

  it('should display message saying app works', () => {
    page.navigateTo();
    expect(page.getParagraphText()).toEqual('app works!');
  });
});
