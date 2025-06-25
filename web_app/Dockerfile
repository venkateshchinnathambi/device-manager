# syntax=docker/dockerfile:1

# ---- BASE IMAGE ----
ARG RUBY_VERSION=3.2.3
FROM ruby:$RUBY_VERSION-slim AS base

WORKDIR /rails

# Install base runtime dependencies
RUN apt-get update -qq && \
    apt-get install --no-install-recommends -y \
    curl \
    libjemalloc2 \
    sqlite3 \
    libpq5 && \
    rm -rf /var/lib/apt/lists/* /var/cache/apt/archives

# Set environment for production
ENV RAILS_ENV="production" \
    BUNDLE_DEPLOYMENT="1" \
    BUNDLE_PATH="/usr/local/bundle" \
    BUNDLE_WITHOUT="development:test"

# ---- BUILD IMAGE ----
FROM base AS build

# Install build-time dependencies
RUN apt-get update -qq && apt-get install --no-install-recommends -y \
  build-essential \
  libpq-dev \
  postgresql-client \
  libyaml-dev \
  pkg-config \
  nodejs \
  yarn \
  git && \
  rm -rf /var/lib/apt/lists/* /var/cache/apt/archives

# Install application gems
COPY Gemfile Gemfile.lock ./
RUN bundle install && \
    rm -rf ~/.bundle /usr/local/bundle/ruby/*/cache /usr/local/bundle/ruby/*/bundler/gems/*/.git

# Copy application code
COPY . .

# Precompile assets
RUN SECRET_KEY_BASE_DUMMY=1 ./bin/rails assets:precompile

# ---- FINAL RUNTIME IMAGE ----
FROM base

# Copy app + gems from build stage
COPY --from=build /usr/local/bundle /usr/local/bundle
COPY --from=build /rails /rails

# Use non-root user for better security
RUN groupadd --system --gid 1000 rails && \
    useradd rails --uid 1000 --gid 1000 --create-home --shell /bin/bash && \
    chown -R rails:rails db log storage tmp

#USER rails:rails

# Entrypoint (optional) to handle db setup/migrations
ENTRYPOINT ["/rails/bin/docker-entrypoint"]


EXPOSE 50051  
# Start server (overridable)
EXPOSE 80
CMD ["./bin/thrust", "./bin/rails", "server"]
