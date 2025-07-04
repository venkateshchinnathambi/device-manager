# syntax=docker/dockerfile:1

# ---- BASE IMAGE ----
ARG RUBY_VERSION=3.2.3
#FROM ruby:$RUBY_VERSION-slim AS base
FROM ruby:3.2.3 AS base


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
  curl \
  gnupg \
  build-essential \
  libpq-dev \
  postgresql-client \
  libyaml-dev \
  pkg-config \
  git && \
  curl -fsSL https://deb.nodesource.com/setup_18.x | bash - && \
  curl -sS https://dl.yarnpkg.com/debian/pubkey.gpg | apt-key add - && \
  echo "deb https://dl.yarnpkg.com/debian/ stable main" | tee /etc/apt/sources.list.d/yarn.list && \
  apt-get update -qq && \
  apt-get install --no-install-recommends -y nodejs yarn && \
  rm -rf /var/lib/apt/lists/* /var/cache/apt/archives


# Install Bundler version to match Gemfile.lock
RUN gem install bundler:2.6.8

# Install application gems
COPY Gemfile Gemfile.lock ./
#RUN bundle _2.6.8_ install && \
#    rm -rf ~/.bundle /usr/local/bundle/ruby/*/cache /usr/local/bundle/ruby/*/bundler/gems/*/.git

RUN gem install bundler -v 2.6.8 && bundle install --verbose




# Copy application code
COPY . .

# Precompile assets
#Hardcode SECRET_KEY_BASE directly
ENV SECRET_KEY_BASE=a7aa7ea2608386f9d49b05e7062667b3c745574d8e04d7393b689a8f48100f42d1e212bf8690543c9a06d7c33c9e307fc8a1373fe99729cd1cfb5d98008499e9
ENV RAILS_ENV=production
ENV RAILS_LOG_LEVEL=debug
RUN ./bin/rails assets:precompile --trace

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
