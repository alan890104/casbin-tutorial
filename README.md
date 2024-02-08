# Casbin Tutorial

This system is designed to manage access control using a combination of roles, permissions, and domains (sites) to ensure that users can only access resources they are authorized to. Let's break down the components and the design for easier understanding:

### System Components

- **Request Definition (r)**: Represents an access request by a user. It includes the user making the request (`sub`), the domain or site they're accessing (`dom`), the object or resource they're trying to access (`obj`), and the action they're trying to perform on that object (`act`).

- **Policy Definition (p)**: Defines the permissions. It specifies which role (`sub`), in which domain (`dom`), can access which resource (`obj`), and what actions (`act`) they can perform.

- **Role Definition (g)**: Defines relationships between entities. It can map a business entity to a user, a user to a role, or a role to a static role, indicating a hierarchy or affiliation.

- **Policy Effect (e)**: Determines the effect of the policies. The system allows an action if there is at least one "allow" policy and no "deny" policy for the request.

- **Matchers (m)**: Specifies the conditions under which a policy applies to a request. It checks if the user has a role that matches the policy, if the domain of the request matches the policy's domain, and if the requested object and action match the policy's object and action through pattern matching.

### System Design

1. **Permissions**: Permissions are specific actions allowed on certain resources within a domain. For example, `cam-admin` can perform any action (`*`) on all cloud camera resources (`/cloud-cameras/*`) within `site1`.

2. **Roles and Role Inheritance**: Users are assigned roles that determine their permissions. Roles can inherit permissions from other roles, enabling flexible and hierarchical permission management. For instance, a `customer(site1)` can inherit permissions from `cam-creator(site1)` and `cam-viewer(site1)`, aggregating their permissions.

3. **Domains/Sites**: The system supports multi-tenancy or domain-specific roles and permissions, allowing distinct access controls for different sites or domains. A user's permissions in `site1` can be different from their permissions in `site2`.

4. **User Assignments**: Users are assigned roles directly or through inheritance. For example, `alice` is assigned the `customer(site1)` role, inheriting the permissions associated with that role within `site1`.

5. **Business Entity and Headquarter Roles**: The system supports higher-level organizational roles, allowing for complex hierarchies. Users or entities with broader responsibilities can be assigned roles like `distributor` or `headquarter`, which can inherit or grant access across multiple sites or domains.

### How It Works

When a user attempts to access a resource, the system checks their request against the defined policies to determine if they have the required permissions. It considers the user's roles, the domain of the request, and the specific resource and action they are attempting to access. If the conditions are met, the access is granted.

This design allows for granular, flexible, and hierarchical access control, suitable for complex organizational structures and multi-tenant environments.
