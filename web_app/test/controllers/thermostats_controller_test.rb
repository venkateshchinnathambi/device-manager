require "test_helper"

class ThermostatsControllerTest < ActionDispatch::IntegrationTest
  test "should get index" do
    get thermostats_index_url
    assert_response :success
  end
end
